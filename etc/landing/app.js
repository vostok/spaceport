function checkConnection() {

    $('#services a').each(function(index) {

        var url = $(this).attr('href');

        $.ajax({
            url: url,
            type: "get",
            cache: false,
            dataType: 'jsonp', // it is for supporting crossdomain
            crossDomain : true,
            asynchronous : false,
            jsonpCallback: 'deadCode',
            timeout : 1500, // set a timeout in milliseconds
            complete : function(xhr, responseText, thrownError) {
                var el = $("#services div").get(index);
                if(xhr.status == "200") {
                    $(el).removeClass("error").addClass("ok");
                }
                else {
                    $(el).removeClass("ok").addClass("error");
                }
            }
        });

    });

    window.setTimeout(checkConnection, 3000);

}
$(document).ready( function() {
    jQuery.migrateMute = true;
    window.setTimeout(checkConnection);
});