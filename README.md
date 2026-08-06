# learning-kubebuilder

## Context

- Want to learn about Kubebuilder, Custom Resource Definition, and the controller logic
  that handles it
- Going to try using the `/teach` skill to learn about it

## Lesson 2

After running Step 7 of Lesson 2, i got

```

2026-08-06T12:02:14-04:00       INFO    setup   Starting manager
2026-08-06T12:02:14-04:00       INFO    starting server {"name": "health probe", "addr": "[::]:8081"}
2026-08-06T12:02:14-04:00       INFO    Starting EventSource    {"controller": "greeting", "controllerGroup": "apps.kubebuilder-lessons.dev", "controllerKind": "Greeting", "source": "kind source: *v1.Greeting"}
2026-08-06T12:02:14-04:00       INFO    Starting Controller     {"controller": "greeting", "controllerGroup": "apps.kubebuilder-lessons.dev", "controllerKind": "Greeting"}
2026-08-06T12:02:14-04:00       INFO    Starting workers        {"controller": "greeting", "controllerGroup": "apps.kubebuilder-lessons.dev", "controllerKind": "Greeting", "worker count": 1}
2026-08-06T12:03:17-04:00       INFO    reconciled greeting     {"controller": "greeting", "controllerGroup": "apps.kubebuilder-lessons.dev", "controllerKind": "Greeting", "Greeting": {"name":"greeting-sample","namespace":"default"}, "namespace": "default", "name": "greeting-sample", "reconcileID": "99f717ce-3e17-40a2-9c3d-d0314bb47c00", "configmap-op": "created"}
2026-08-06T12:04:34-04:00       INFO    reconciled greeting     {"controller": "greeting", "controllerGroup": "apps.kubebuilder-lessons.dev", "controllerKind": "Greeting", "Greeting": {"name":"greeting-sample","namespace":"default"}, "namespace": "default", "name": "greeting-sample", "reconcileID": "e197dae5-5942-41b0-8f49-6e958716fb81", "configmap-op": "updated"}
```

Indicating the reconcile loop actually worked, it took the spec value and then created a config map operation for it

## Resources

- https://book.kubebuilder.io/
- https://www.aihero.dev/skills-teach
