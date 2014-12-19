So {{ .user.Name }} it looks like you're trying to reset your password.  Follow this link within 24 hours and we'll reset it for you.

<a href="{{ .request.Host }}/admin/redeem/{{ .reset.Token }}">Reset</a>