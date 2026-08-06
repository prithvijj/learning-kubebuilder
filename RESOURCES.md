# Kubebuilder & Controller Internals Resources

## Knowledge

- [The Kubebuilder Book](https://book.kubebuilder.io/)
  The canonical, official guide from kubernetes-sigs. Use for: scaffolding workflow, project layout, CRD/webhook how-tos. Start here for anything hands-on.
- [Kubebuilder Book — Architecture Concepts](https://book.kubebuilder.io/architecture.html)
  Official diagram + explanation of Manager, Cache, Client, Controller, Reconciler and how they wire together. Use for: the mental model of controller-runtime's pieces.
- [kubernetes/sample-controller — controller-client-go.md](https://github.com/kubernetes/sample-controller/blob/master/docs/controller-client-go.md)
  The canonical explainer of the client-go informer/workqueue pattern that controller-runtime is itself built on top of. Use for: understanding informers, listers, and workqueues at the level below Kubebuilder's abstractions.
- [controller-runtime `reconcile` package docs (pkg.go.dev)](https://pkg.go.dev/sigs.k8s.io/controller-runtime/pkg/reconcile)
  Primary source for the exact `Reconcile(ctx, Request) (Result, error)` contract. Use for: precise semantics of return values, requeueing, errors.
- [Kubernetes Blog — "How the controller-runtime Cache Actually Works" (2026)](https://kubernetes.io/blog/2026/07/29/controller-runtime-cache-explained/)
  Recent official k8s.io blog post on the cache/informer layer and why controllers don't hammer the API server. Use for: cache internals, resync behavior, list+watch mechanics.
- [Kubernetes docs — Custom Resources concept page](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/)
  Official explanation of what a CRD/CR actually is at the API machinery level (as opposed to what a controller does with it). Use for: grounding "a CR is just stored state" before introducing controllers.
- [_Kubernetes Patterns_ (2nd ed.) — Ibryam & Huß, O'Reilly](https://www.oreilly.com/library/view/kubernetes-patterns-2nd/9781098131678/)
  Especially the Controller and Operator chapters. Use for: design-level patterns once the mechanics are solid — good for the "design one from scratch" success criterion.
- [Kubebuilder Book — CronJob Tutorial](https://book.kubebuilder.io/cronjob-tutorial/cronjob-tutorial)
  The official end-to-end worked example (a full CRD + controller, more involved than our toy `Greeting`). Use for: the next step up once the basics from Lesson 2 feel easy, especially for status conditions and owned-object watches.
- [k3d documentation](https://k3d.io/)
  Official docs for the cluster tool the user tests with. Use for: cluster lifecycle (create/delete), image import for locally-built controller images, kubeconfig merging behavior.

## Wisdom (Communities)

- [Kubernetes Slack](https://kubernetes.slack.com/) — `#kubebuilder` and `#sig-api-machinery` channels
  Where kubebuilder/controller-runtime maintainers and heavy users hang out. Use for: design review, "is this the idiomatic way" questions.
- [r/kubernetes](https://reddit.com/r/kubernetes)
  General but active; good for war stories about running real operators in production.

## Gaps
- No resource yet specifically on CRD versioning/conversion webhooks — add when we get there.
- No resource yet on testing controllers (envtest) — add before the first hands-on controller lesson needs it.
