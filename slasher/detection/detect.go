package detection

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	ethpb "github.com/prysmaticlabs/ethereumapis/eth/v1alpha1"
	"github.com/prysmaticlabs/prysm/shared/sliceutil"
	"github.com/prysmaticlabs/prysm/slasher/detection/attestations/types"
)

func (ds *Service) detectAttesterSlashings(
	ctx context.Context,
	att *ethpb.IndexedAttestation,
) ([]*ethpb.AttesterSlashing, error) {
	slashings := make([]*ethpb.AttesterSlashing, 0)
	for i := 0; i < len(att.AttestingIndices); i++ {
		valIdx := att.AttestingIndices[i]
		surroundedAttSlashings, err := ds.detectSurroundVotes(ctx, valIdx, att)
		if err != nil {
			return nil, errors.Wrap(err, "could not detect surround votes on attestation")
		}
		doubleAttSlashings, err := ds.detectDoubleVotes(ctx, valIdx, att)
		if err != nil {
			return nil, errors.Wrap(err, "could not detect double votes on attestation")
		}
		if len(surroundedAttSlashings) > 0 {
			log.Infof("Found %d slashings for val idx %d", len(surroundedAttSlashings), valIdx)
		}
		newSlashings := append(surroundedAttSlashings, doubleAttSlashings...)
		slashings = append(slashings, newSlashings...)
	}
	return slashings, nil
}

func alertOwner(mode string, message string) {
	url := fmt.Sprintf("http://localhost:5000/alert/%s", mode)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(message)))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to alert owner")
		panic("Failed to alert owner")
	}
	defer resp.Body.Close()
}

var attestedSlots = map[uint64]string{}
var alerted = map[uint64]bool{}

// detectDoubleVote --
// TODO(#4589): Implement.
func (ds *Service) detectDoubleVotes(
	ctx context.Context,
	validatorIdx uint64,
	att *ethpb.IndexedAttestation,
) ([]*ethpb.AttesterSlashing, error) {

	validator, err := ds.beaconClient.GetValidator(ctx, validatorIdx)

	hexValidatorPubkey := make([]byte, hex.EncodedLen(len(validator.PublicKey)))
	hex.Encode(hexValidatorPubkey, validator.PublicKey)

	fmt.Printf("%s\n", hexValidatorPubkey)
	if err != nil {
		panic(err)
	}

	slot := att.Data.Slot

	if prevdata, exists := attestedSlots[slot]; exists {
		if att.Data.String() != prevdata {
			// alerted[slot] = true

			if isAlerted := alerted[slot]; !isAlerted {
				alerted[slot] = true
				fmt.Printf("Double vote detected at slot %d. Alert owner\n", slot)
				msg := fmt.Sprintf("{\"slot\": %d, \"pubkey\": \"%s\"}", slot, hexValidatorPubkey)
				alertOwner("telegram", msg)
			}
		}
	} else {
		attestedSlots[att.Data.Slot] = att.Data.String()
	}

	return nil, nil
}

// detectSurroundVotes cross references the passed in attestation with the requested validator's
// voting history in order to detect any possible surround votes.
func (ds *Service) detectSurroundVotes(
	ctx context.Context,
	validatorIdx uint64,
	incomingAtt *ethpb.IndexedAttestation,
) ([]*ethpb.AttesterSlashing, error) {
	res, err := ds.minMaxSpanDetector.DetectSlashingForValidator(
		ctx,
		validatorIdx,
		incomingAtt.Data,
	)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	if res.Kind != types.SurroundVote {
		return nil, nil
	}

	var slashings []*ethpb.AttesterSlashing
	otherAtts, err := ds.slasherDB.IndexedAttestationsForEpoch(ctx, res.SlashableEpoch)
	if err != nil {
		return nil, err
	}
	for _, att := range otherAtts {
		if att.Data == nil {
			continue
		}
		// If there are no shared indices, there is no validator to slash.
		if len(sliceutil.IntersectionUint64(att.AttestingIndices, incomingAtt.AttestingIndices)) == 0 {
			continue
		}

		if isSurrounding(incomingAtt, att) || isSurrounded(incomingAtt, att) {
			slashings = append(slashings, &ethpb.AttesterSlashing{
				Attestation_1: att,
				Attestation_2: incomingAtt,
			})
		}
	}
	if len(slashings) == 0 {
		return nil, errors.New("unexpected false positive in surround vote detection")
	}
	return slashings, nil
}

func isSurrounding(incomingAtt *ethpb.IndexedAttestation, prevAtt *ethpb.IndexedAttestation) bool {
	return incomingAtt.Data.Source.Epoch < prevAtt.Data.Source.Epoch &&
		incomingAtt.Data.Target.Epoch > prevAtt.Data.Target.Epoch
}

func isSurrounded(incomingAtt *ethpb.IndexedAttestation, prevAtt *ethpb.IndexedAttestation) bool {
	return incomingAtt.Data.Source.Epoch > prevAtt.Data.Source.Epoch &&
		incomingAtt.Data.Target.Epoch < prevAtt.Data.Target.Epoch
}
