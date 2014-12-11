{{ define "Content" }}
    <div class="row">
        <div class="col-lg-12">
            <h1 class="page-header">
                Dashboard <small>Statistics Overview</small>
            </h1>
        </div>
    </div>
    <!-- /.row -->

    <div class="row">
        <div class="col-lg-4">
            <table class="table table-bordered table-hover table-striped">
                <thead>
                    <tr>
                        <th>Title</th>
                        <th>Posts</th>
                    </tr>
                </thead>
                <tbody>
                    {{ range .cats }}
                        <tr>
                            <td>{{ .Title }}</td>
                            <td>{{ .Count }}</td>
                        </tr>
                    {{ end }}
                </tbody>
            </table>
        </div>
        <div class="col-lg-8">
            <table class="table table-bordered table-hover table-striped">
                <thead>
                    <tr>
                        <th>Title</th>
                        <th>Author</th>
                        <th>Created</th>
                        <th>Published</th>
                        <th>Edit</th>
                    </tr>
                </thead>
                <tbody>
                    {{ range .recent }}
                        <tr>
                            <td>{{ .Title }}</td>
                            <td>{{ .Author.Name }}</td>
                            <td>{{ fmt_date .Created }}</td>
                            <td>{{ fmt_bool .Published }}</td>
                            <td><a href="/admin/post/{{ .Id }}/edit">Edit</a></td>
                        </tr>
                    {{ end }}
                </tbody>
            </table>
        </div>
    </div>
{{ end }}

{{ define "Scripts" }}
<!-- Morris Charts JavaScript -->
    <script src="/admin/js/plugins/morris/raphael.min.js"></script>
    <script src="/admin/js/plugins/morris/morris.min.js"></script>
    <script src="/admin/js/plugins/morris/morris-data.js"></script>
{{ end }}