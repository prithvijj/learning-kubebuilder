# cert-manager: surface-level understanding only, stopped before deep reading

Ended Lesson 7 with "a basic understanding of what those controllers are" — recognizes the multi-controller-per-CRD shape and that ProcessItem/Reconcile are the same contract in different clothes, but did not do the directed line-by-line reading exercise (Step 1's three-item scavenger hunt) or report back findings. Treat cert-manager specifics (trigger/keymanager/requestmanager/issuing/readiness/revisionmanager split, the ACME finalizer) as introduced but not demonstrated — don't assume recall of exact function/condition names without a refresher.

## Implications
- If a future session returns to cert-manager, start with a light recap rather than assuming the Step 1/2 details stuck.
- The broader mission ([[MISSION.md]]) success criteria are otherwise all exercised at least once (see [[0005-envtest-suite-complete]]); this was the "breadth" extension, not a core gap.
