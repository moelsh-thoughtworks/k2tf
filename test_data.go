package main

const configMapYAML = `
apiVersion: v1
kind: ConfigMap
metadata:
  name: fooConfigMap
  namespace: bar
  labels:
    lbl1: somevalue
    lbl2: another
data:
  item1: wow
  item2: wee
`

const configMapHCL = `resource "kubernetes_config_map" "foo" {
  metadata {
    name      = "fooConfigMap"
    namespace = "bar"
    labels {
      lbl1 = "somevalue"
      lbl2 = "another"

    }

  }
  data {
    item1 = "wow"
    item2 = "wee"

  }
}`

const basicDeploymentYAML = `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: baz-app
  namespace: bat
  annotations:
    foo: fam
spec:
  replicas: 2
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      annotations:
        foo: fam
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx
        args: ["--debug", "--test"]
        ports:
        - port: 80
          containerPort: 80
        resources:
          requests:
            cpu: "1"
            memory: "512Mi"
          limits:
            cpu: "1"
            memory: "1Gi"
`

const basicDeploymentHCL = `resource "kubernetes_deployment" "foo" {
  metadata {
    name      = "baz-app"
    namespace = "bat"
    annotations {
      foo = "fam"

    }

  }
  spec {
    replicas = 2
    selector {
      match_labels {
        app = "nginx"

      }

    }
    template {
      metadata {
        labels {
          app = "nginx"

        }
        annotations {
          foo = "fam"

        }

      }
      spec {
        container {
          name  = "nginx"
          image = "nginx"
          args  = ["--debug", "--test"]
          port {
            container_port = 80

          }
          resources {
            limits {
              cpu    = "1"
              memory = "1Gi"

            }
            requests {
              cpu    = "1"
              memory = "512Mi"

            }

          }

        }

      }

    }

  }
}
`

const prometheus_pod = `apiVersion: v1
kind: Pod
metadata:
  annotations:
    prometheus.io/port: "9100"
    prometheus.io/scheme: http
    prometheus.io/scrape: "true"
  creationTimestamp: 2018-09-10T23:29:40Z
  generateName: node-exporter-
  labels:
    controller-revision-hash: "2418008739"
    name: node-exporter
    pod-template-generation: "1"
  name: node-exporter-7fth7
  namespace: prometheus
  ownerReferences:
  - apiVersion: apps/v1
    blockOwnerDeletion: true
    controller: true
    kind: DaemonSet
    name: node-exporter
    uid: 08e57633-594e-11e8-bdee-42010a800279
  resourceVersion: "42836226"
  selfLink: /api/v1/namespaces/prometheus/pods/node-exporter-7fth7
  uid: 63a4e2b5-b551-11e8-ab49-42010a8000db
spec:
  automountServiceAccountToken: true
  containers:
  - image: prom/node-exporter
    imagePullPolicy: Always
    name: prom-node-exporter
    ports:
    - containerPort: 9100
      name: metrics
      protocol: TCP
    resources: {}
    securityContext:
      privileged: true
      readOnlyRootFilesystem: false
      runAsNonRoot: false
      runAsUser: 0
    terminationMessagePath: /dev/termination-log
    terminationMessagePolicy: File
    volumeMounts:
    - mountPath: /var/run/secrets/kubernetes.io/serviceaccount
      name: default-token-rkd4g
      readOnly: true
  dnsPolicy: ClusterFirst
  hostPID: true
  nodeName: gke-cloudlogs-dev-default-pool-4a2a9dae-9b01
  restartPolicy: Always
  schedulerName: default-scheduler
  securityContext: {}
  serviceAccount: default
  serviceAccountName: default
  terminationGracePeriodSeconds: 30
  tolerations:
  - effect: NoExecute
    key: node.kubernetes.io/not-ready
    operator: Exists
  - effect: NoExecute
    key: node.kubernetes.io/unreachable
    operator: Exists
  - effect: NoSchedule
    key: node.kubernetes.io/disk-pressure
    operator: Exists
  - effect: NoSchedule
    key: node.kubernetes.io/memory-pressure
    operator: Exists
  volumes:
  - name: default-token-rkd4g
    secret:
      defaultMode: 420
      secretName: default-token-rkd4g
status:
  conditions:
  - lastProbeTime: null
    lastTransitionTime: 2018-09-10T23:29:43Z
    status: "True"
    type: Initialized
  - lastProbeTime: null
    lastTransitionTime: 2018-09-10T23:30:21Z
    status: "True"
    type: Ready
  - lastProbeTime: null
    lastTransitionTime: 2018-09-10T23:29:43Z
    status: "True"
    type: PodScheduled
  containerStatuses:
  - containerID: docker://320ffc1ab868431b0d51b99d2aac02122b4b03d8274b2c59640e00c7fb6b7a45
    image: prom/node-exporter:latest
    imageID: docker-pullable://prom/node-exporter@sha256:55302581333c43d540db0e144cf9e7735423117a733cdec27716d87254221086
    lastState: {}
    name: prom-node-exporter
    ready: true
    restartCount: 0
    state:
      running:
        startedAt: 2018-09-10T23:30:21Z
  hostIP: 172.16.0.3
  phase: Running
  podIP: 10.8.1.3
  qosClass: BestEffort
  startTime: 2018-09-10T23:29:43Z`

const volumesHCL = `resource "kubernetes_pod" "foo" {
  spec {
    volume {
      name = "default-token-rkd4g"
        secret {
          secret_name  = "default-token-rkd4g"
          default_mode = 420

        }

      }

    }
  }
}
`

const volumesYAML = `apiVersion: v1
kind: Pod
spec:
  volumes:
  - name: default-token-rkd4g
    secret:
      defaultMode: 420
      secretName: default-token-rkd4g
`

// - name: some-volume
//   configMap:
//     defaultMode: 420
//     name: cm1
