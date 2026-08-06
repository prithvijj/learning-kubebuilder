# First working controller: CR edit → Reconcile → ConfigMap update

Completed Lesson 2 end-to-end: scaffolded `greeting-operator` with Kubebuilder, ran it against a k3d cluster with `make run`, and confirmed that editing a `Greeting` CR's `spec.message` and re-applying it caused `Reconcile` to fire again and the derived ConfigMap to update in place (not duplicate). This is direct evidence of understanding the level-triggered/idempotent reconcile pattern from [[0001-anatomy-of-a-controller]] in practice, not just in the abstract diagram — safe to build on this without re-explaining "why does the same function handle create and update."

Also hit and fixed two real scaffolding errors independently after a hint (missing `apierrors` import, and a `resourcees` typo in an RBAC marker breaking `controller-gen`) — comfortable reading Go compiler and `controller-gen` marker-parser errors and mapping them back to source lines.

## Implications
- Ready for the next layer: currently the controller only watches `Greeting` itself, not the ConfigMap it creates — the `Owns()` watch and status reporting are both open per [[0002-scaffold-your-first-controller]]'s closing choice.
