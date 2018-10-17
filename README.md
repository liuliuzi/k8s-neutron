# k8s-neutron
Modify k8s network dynamically

Deploy:
            K8s worknode                             K8s master node

+-------------------------------+   +-----------------------------------------------+
|                   cni plugin  |   |      +----------------+ +------------------+  |
|     +--------------+          |   |      | kube-apiserver | |  kube-scheduler  |  |
|     |  kubelet     |          |   |      |                | |                  |  |
|     +--------------+          |   |      +----------------+ +------------------+  |
|     +--------------+          |   |      +---------------------------------+      |
|     | kube-proxy   |          |   |      |  kube-controller-manger         |      |
|     |              |          |   |      |                                 |      |
|     +--------------+          |   |      +---------------------------------+      |
|     +--------------------+    |   |      +---------------------------------+      |
|     | kube-neutron-agent |    |   |      |  kube-neutron-server            |      |
|     |                    |    |   |      |                                 |      |
|     +--------------------+    |   |      +---------------------------------+      |
+-------------------------------+   +-----------------------------------------------+









