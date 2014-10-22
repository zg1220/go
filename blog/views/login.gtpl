<html>
 <head>
 <script type="text/javascript" src="../static/js/jquery-2.1.1.js"></script>

 <script type="text/javascript">
 $(document).ready(function(){
 	$("form").submit(function(){
 		id=$("#login_id").val();
 		pwd=$("#login_pwd").val();		
 		if(id=="" || pwd==""){
 			alert("nil");
 			return false;
 		}
 		else{
 			$.ajax({
 				type:"GET",
 				url:"ajax?",//login="+id,
 				data:{"login":id,"pwd":pwd},
 				success:function(msg){
 					if (msg=="1"){
 						window.location.href="/success";
 					}
 					else{alert("error!");}
 				}
 			});
 			return false;
 		};
 	});
});
</script>
 </head>
 <body>
  <form method="post" action="#" autocomplete="off">
   <input type="text" name="login" id="login_id" >
   <input type="password" name="pwd" id="login_pwd">
   <button name="btn" type="submit" >text</button>
  </form>
  <P><span id="txtHint"></span></P>
 </body>
</html>