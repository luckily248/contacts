 				<tr>
                 <form class="form ajax" id="add" action="/contacts/{{.GetId}}/update" method="post" data-refresh="#contacts">
				<td></td>
                  <td><div class="name"><input class="form-control" name="name" type="text" value={{.Name}}></input></div></td>
                  <td><div class="remark"><input class="form-control" name="remark" type="text" value={{.Remark}}></input></div></td>
                  <td><div class="phone1"><input class="form-control" name="phone1" type="text" value={{.Phone1}}></input></div></td>
				  <td><div class="phone2"><input name="phone2" type="text" value={{.Phone2}}></input></div></td>
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