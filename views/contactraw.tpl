 				<tr>
                 <form class="form ajax" id="add" action="/contacts/{{.GetId}}/update" method="post" data-refresh="#contacts">
				<td></td>
                  <td><input name="name" type="text" value={{.Name}}></input></td>
                  <td><input name="remark" type="text" value={{.Remark}}></input></td>
                  <td><input name="phone" type="text" value={{.Phone}}></input></td>
					<td>
					
				<button  class="btn btn-success ajax" type="submit" >
					<i class="glyphicon glyphicon-ok">
					</i>
					</button>
					<a  class="btn btn-danger 	ajax" data-refresh="#contacts">
					<i class="glyphicon glyphicon-remove">
					</i>
					</a>
					</td>
				</form>
                </tr>