{{ define "Title" }}
Login
{{ end }}

{{ define "Content" }}
<form role="form" method="POST" action="/login">
    <div class="form-group">
        <label for="emailInput">Email address</label>
        <input type="email" name="email" class="form-control" id="emailInput" placeholder="Enter Email">
    </div>
    <div class="form-group">
        <label for="passwordInput">Password</label>
        <input type="password" name="password" class="form-control" id="passwordInput" placeholder="Enter Password">
    </div>
    <button type="submit" class="btn btn-primary">Submit</button>
</form>
<a href="/admin/forgot">Forgot Password?</a>
{{ end }}