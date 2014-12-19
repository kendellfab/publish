{{ define "Title" }}
Reset Password &middot; {{ .user.Name }}
{{ end }}

{{ define "Content" }}
<form role="form" method="POST" action="/admin/redeem">
	<input type="hidden" name="user_id" value="{{ .user.Id }}" />
	<input type="hidden" name="token" value="{{ .reset.Token }}" />
    <div class="form-group">
        <label for="passwordInput">Password</label>
        <input type="password" name="password" class="form-control" id="passwordInput" placeholder="Enter Password">
    </div>
    <button type="submit" class="btn btn-primary">Submit</button>
</form>
{{ end }}