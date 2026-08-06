# Prior knowledge baseline: Go + Kubernetes The Hard Way

Established at mission intake, not yet demonstrated in a lesson: strong Go experience (concurrency, interfaces), and has completed _Kubernetes The Hard Way_, so already understands API server / etcd / kubelet / scheduler and how the base control plane fits together. This means lessons can skip Go syntax entirely and skip re-explaining what the API server or etcd are — start from "a CRD registers a new type with the API server" rather than from "what is the API server."

## Implications
- Safe to move fast through anything that's pure Go.
- Not yet established: any familiarity with client-go, informers, workqueues, or controller-runtime specifically — treat these as genuinely new (see [[MISSION.md]]).
