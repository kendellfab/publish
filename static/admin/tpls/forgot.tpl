{{ define "Title" }}
Forgot Password
{{ end }}

{{ define "Content" }}
<form role="form" method="POST" action="/admin/forgot">
    <div class="form-group">
        <label for="emailInput">Email address</label>
        <input type="email" name="email" class="form-control" id="emailInput" placeholder="Enter Email">
    </div>
    <button type="submit" class="btn btn-primary">Submit</button>
</form>
{{ end }}