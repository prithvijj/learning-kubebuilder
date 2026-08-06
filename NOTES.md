# Notes

- User has strong Go and completed _Kubernetes The Hard Way_ — don't re-teach Go syntax or base control-plane concepts (API server/etcd/kubelet/scheduler). Build from there. See [[MISSION.md]].
- No stated pacing/scope constraints as of 2026-08-06 — free to go broad (webhooks, versioning, etc.) as it becomes relevant.
- User tests with **k3d**, not kind or a managed cluster. Use k3d in every hands-on lesson from Lesson 2 onward. As of 2026-08-06 they already have a cluster named `mycluster` (k3d v5.9.0, k3s v1.35.5) reachable via context `k3d-mycluster`.
- Environment observed 2026-08-06: Go 1.25.6, kubectl v1.36.2, Docker 29.5.3 all present; `kubebuilder` CLI is NOT installed yet — install it in Lesson 2.
