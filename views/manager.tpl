<!DOCTYPE html>

<html>
<head>
   <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <!-- 上述3个meta标签*必须*放在最前面，任何其他内容都*必须*跟随其后！ -->
    <meta name="description" content="">
    <meta name="author" content="">
    <link rel="icon" href="../../favicon.ico">

    <title>企业通讯录</title>

    <!-- Bootstrap core CSS -->
    <link href="static/css/bootstrap.min.css" rel="stylesheet">

    <!-- Custom styles for this template -->
    <!--<link href="dashboard.css" rel="stylesheet">-->
</head>

<body>
<h2 class="sub-header">企业通讯录</h2>
  <div class="table-responsive">
            <table class="table table-striped">
              <thead>
                <tr>
                  <th>编号</th>
                  <th>名字</th>
                  <th>职位</th>
                  <th>电话</th>
                </tr>
              </thead>
              
				<tbody id="contacts" data-refresh-url="/contacts/getalltable">
					{{range $k,$v:=.contacts}}
				
           		     <tr>
           		       <td>{{$k}}</td>
           		       <td>{{$v.Name}}</td>
           		       <td>{{$v.Remark}}</td>
           		       <td>{{$v.Phone}}</td>
						<td>
					
				<a  class="btn ajax" href="/contacts/{{$v.GetId}}/preupdate" data-method="post" data-replace-closest="tr">
					<i class="glyphicon glyphicon-pencil">
					</i>
					</a>
		 	  	<a  class="btn ajax" href="/contacts/{{$v.GetId}}/del" data-method="post" data-remove-closest="tr">
					<i class="glyphicon glyphicon-trash">
					</i>
					</a>
					</td>
                </tr>
				{{end}}
				 </tbody>
			
             
            </table>
			<a class="btn ajax" href="/contacts/add" data-method="post" data-append="#contacts">添加联系人</a>
          </div>
	<script>
	)

	</script>
<!-- jQuery (necessary for Bootstrap's JavaScript plugins) -->
    <script src="static/js/jquery.min.js"></script>
    <!-- Include all compiled plugins (below), or include individual files as needed -->
    <script src="static/js/bootstrap.min.js"></script>
	<script src="static/js/eldarion-ajax.min.js"></script>
</body>

</html>
