# Finalizers block the API server's delete, not the controller

Initial phrasing was "finalizers block the controller from deleting the object." Corrected: the controller never deletes anything here — it's the API server that's blocked from removing the object from etcd while a finalizer remains. The controller keeps reconciling normally throughout; a stuck finalizer means the *object* is stuck, not the controller.

## Implications
- Worth double-checking this distinction stays clear when webhooks come up later — admission webhooks intercept the API server's write path directly, a genuinely different insertion point, and conflating "controller does X" with "API server does X" would compound here.
