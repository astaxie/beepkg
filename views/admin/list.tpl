<div class="row-fluid">
                    
<div class="btn-toolbar">
    <a class="btn btn-primary" href="/admin/addpkg"><i class="icon-plus"></i> 添加PKG</a>
  <div class="btn-group">
  </div>
</div>
<div class="well">
    <table class="table">
      <thead>
        <tr>
          <th>#</th>
          <th>PkgName</th>
          <th style="width: 26px;"></th>
        </tr>
      </thead>
      <tbody>
        {{range .Articles}}
        <tr>
          <td>{{.Id}}</td>
          <td>{{.Content}}</td>
          <td>
              <a href="/admin/editpkg/{{.Id}}"><i class="icon-pencil"></i></a>
              <a href="/admin/delpkg/{{.Id}}" role="button" data-toggle="modal"><i class="icon-remove"></i></a>
          </td>
        </tr>
        {{end}}
      </tbody>
    </table>
</div>

<!--
<div class="pagination">
    <ul>
        <li class="active"><a href="#">上一页</a></li>
        <li><a href="/admin/index?page=1">1</a></li>
        <li><a href="#">2</a></li>
        <li><a href="#">下一页</a></li>
    </ul>
</div>

-->