$(function(){
    $.ajax({
        url:'/blogList',
        type:'POST',
        success:function (data) {
            var filenames = data.files;
            var fileArray = filenames.split(";");
            for(var i=0;i<fileArray.length;i++){
                var fileName = fileArray[i].substring(fileArray[i].lastIndexOf("\/") + 1)
                var html = "<li class=\"list-group-item\"><a href='javascript: void(0);' onclick='modify(\'"+fileName+"\')'>"+fileName+"</a></li>";
                var fileItem = $(html);
                $("#fileList").append(fileItem);
            }
        }
    })
})

function modify(filename) {
    $.ajax({
        url:'/createNew?name=' + fileName,
        type:'GET',
        success:function (e) {

        }
    })
}