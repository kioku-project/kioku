{{/*
Expand the name of the chart.
*/}}
{{- define "kioku.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "kioku.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "kioku.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "kioku.labels" -}}
helm.sh/chart: {{ include "kioku.chart" . }}
{{ include "kioku.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}



{{/*
Database labels
*/}}
{{- define "kioku.database.labels" -}}
app.kubernetes.io/name: {{ include "kioku.fullname" . }}-database
{{- end }}


{{/*
Frontend labels
*/}}
{{- define "kioku.frontend.labels" -}}
app.kubernetes.io/name: {{ .Values.frontend.name }}
{{- end }}


{{/*
Login labels
*/}}
{{- define "kioku.carddeck.labels" -}}
app.kubernetes.io/name: {{ .Values.carddeck.name }}
{{- end }}

{{/*
Srs labels
*/}}
{{- define "kioku.srs.labels" -}}
app.kubernetes.io/name: {{ .Values.srs.name }}
{{- end }}

{{/*
Register labels
*/}}
{{- define "kioku.user.labels" -}}
app.kubernetes.io/name: {{ .Values.user.name }}
app: {{ .Values.user.name }}
{{- end }}

{{/*
Frontend Proxy labels
*/}}
{{- define "kioku.frontend_proxy.labels" -}}
app.kubernetes.io/name: {{ .Values.frontend_proxy.name }}
{{- end }}

{{/*
Collaboration labels
*/}}
{{- define "kioku.collaboration.labels" -}}
app.kubernetes.io/name: {{ .Values.collaboration.name }}
{{- end }}

{{/*
Notification labels
*/}}
{{- define "kioku.notification.labels" -}}
app.kubernetes.io/name: {{ .Values.notification.name }}
{{- end }}


{{/*
Selector labels
*/}}
{{- define "kioku.selectorLabels" -}}
app.kubernetes.io/name: {{ include "kioku.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "kioku.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "kioku.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}
