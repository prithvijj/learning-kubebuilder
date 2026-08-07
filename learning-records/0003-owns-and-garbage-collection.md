# Owns() self-healing and the GC/Reconcile distinction

Completed Lesson 3: added `.Owns(&corev1.ConfigMap{})`, confirmed a hand-edited ConfigMap gets reconciled back to match `Greeting.Spec.Message`, and confirmed a deleted ConfigMap gets recreated. Also correctly worked through the trap at the end — deleting the Greeting deletes the ConfigMap via Kubernetes' garbage collector, not via any Reconcile call, since the manager logs show no reconcile firing for that deletion.

## Implications
- Owner references, secondary watches, and GC-driven cascading delete are now solid — safe to reference these without re-explaining in future lessons (e.g. [[0004-reporting-status]] and beyond).
- Good foundation for reading real operators later: this Owns()/GC split is exactly how built-ins like Deployment→ReplicaSet→Pod work.
