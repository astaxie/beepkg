<link rel="stylesheet" type="text/css" href="/static/css/demo.css" />
        
<script type="text/javascript" src="/static/js/Markdown.Converter.js"></script>
<script type="text/javascript" src="/static/js/Markdown.Sanitizer.js"></script>
<script type="text/javascript" src="/static/js/Markdown.Editor.js"></script>
<div class="row-fluid">
<div class="well">
    <div id="myTabContent" class="tab-content">
      <div class="tab-pane active in" id="home">
    <form id="tab" method="post">
        <label>PKG名称</label>
        <input type="text" name="pkgname" value="{{.pkg.Pathname}}" class="input-xlarge" readonly>
		<input type="hidden" name="id" value="{{.article.Id}}">
        <div class="wmd-panel">
            <div id="wmd-button-bar"></div>
            <textarea name="content" class="wmd-input" id="wmd-input">{{.article.Content}}</textarea>
        </div>
        <div id="wmd-preview" class="wmd-panel wmd-preview"></div>
        <div class="btn-toolbar">
        	<button class="btn btn-primary" type="submit"><i class="icon-save"></i> 提交</button>
		</div>
    </form>
      </div>
  </div>

</div>
<script type="text/javascript">
(function () {
      var converter1 = Markdown.getSanitizingConverter();
      var editor1 = new Markdown.Editor(converter1);
      editor1.run();
})();
</script>