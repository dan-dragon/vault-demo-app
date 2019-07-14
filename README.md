## Debug Rendering on correct template format as running sidecar.

Sample Usage:
docker build -t vault-demo-app:0.1.0 .

in deployment.yaml

add sidecar:

      - name: app-local
        image: vault-demo-app:0.1.0
        imagePullPolicy: IfNotPresent
        env:
        - name: SECRETS_FILE
          value: "/etc/secrets/YOUR_SECRET_DESTINATION_FILE"
        volumeMounts:
        - name: secrets
          mountPath: /etc/secrets
