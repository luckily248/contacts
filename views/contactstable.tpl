<tbody id="contacts" data-refresh-url="/contacts/getalltable">
					{{range $k,$v:=.}}
				
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