# Mission: Kubebuilder & Controller Internals

## Why
You have strong Go experience and have already been through _Kubernetes The Hard Way_, so you understand the control plane from the infra side (API server, etcd, scheduler, kubelet). What's missing is the other half: what actually happens when you define your **own** API type (a Custom Resource) and how the logic that reacts to it gets run. You want to open up the Kubebuilder/controller-runtime black box.

## Success looks like
- Explain the controller-runtime machinery end-to-end: from `kubectl apply` of a CR, through the API server, into a watch/informer, through a workqueue, to a `Reconcile` call — using the right vocabulary at each step.
- Scaffold and write a CRD + controller from scratch with Kubebuilder, including the API types and the `Reconcile` function body.
- Read and confidently modify an existing real-world operator codebase (at work or open source).
- Look at a new problem and design a sound CRD + controller architecture for it before writing any code.

## Constraints
- No time/pace constraints stated. No topics explicitly excluded — go broad and deep as it comes up naturally (webhooks, CRD versioning/conversion, etc. are all in scope if relevant).

## Out of scope
- (none specified yet — revisit if the user wants to fence anything off)

## Prior knowledge assumed
- Strong Go (concurrency, interfaces, generics-era Go) — do not re-teach language basics.
- Completed _Kubernetes The Hard Way_ — knows API server, etcd, kubelet, scheduler, and how the base control plane fits together. Build on this instead of re-explaining it.
- New to: client-go/controller-runtime internals (informers, workqueues, reconcilers), Kubebuilder scaffolding/tooling.
