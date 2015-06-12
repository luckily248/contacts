 				<tr>
                 <form class="form ajax" action="/contacts/{{.GetId}}/update" data-method="post">
                  <td><input name="name" type="text">{{.Name}}</input></td>
                  <td><input name="remark" type="text">{{.Remark}}</input></td>
                  <td><input name="phone" type="text">{{.Phone}}</input></td>
					<td>
					
				<a  class="btn ajax" type="submit" >
					<i class="glyphicon glyphicon-ok">
					</i>
					</a>
					<a  class="btn ajax" >
					<i class="glyphicon glyphicon-remove">
					</i>
					</a>
					</td>
					</form>
                </tr>