$(function() {
    let file_list = $('#file_list');
    let cnt = 0;

    $.get('/files', function(file_list) {
        file_list.forEach(e => {
            AddFile(e)
        });
    });

    $.get("/storage", function(storage_size) {
        $("#storage").text("Used Storage: " + storage_size);
    })

    $(document).ready(function() {
        function fetchNodeCnt() {
            $.get('/node_cnt', function(data) {
                $("#node_cnt").text("Alive Node Count : " + data);
            });
        }
        fetchNodeCnt();

        setInterval(fetchNodeCnt, 4000);
    });

    $(document).ready(function() {
        $('#file_upload_button').click(function() {
            let file_data = $('#user_file').prop('files')[0];
            let form_data = new FormData();

            form_data.append('user_file', file_data);
            $.ajax({
                url: '/upload/0',
                dataType: 'json',
                contentType: false,
                processData: false,
                data: form_data,
                type: 'post',
                success: function(response) {
                    AddFile(response);

                    location.reload();
                }
            });
        });
    });
   
    let AddFile = function(file) {
        let row = "<tr><td>" + (++cnt) + "</tb><td>" + file.name + "</td><td>" + file.size + "Bytes" + "</td><td>" + file.create_at + "</td><td><a href='/download/" + file.id + "' download>Download</a></td></tr>";

        $("#file_list tbody").append(row);
    }

    $(document).ready(function() {
        $('#file_upload_all_button').click(function() {
            let file_data = $('#user_file').prop('files')[0];
            let form_data = new FormData();
    
            form_data.append('user_file', file_data);
            $.ajax({
                url: '/upload/1',
                dataType: 'json',
                contentType: false,
                processData: false,
                data: form_data,
                type: 'post',
                success: function() {
                    location.reload();
                }
            });
        });
    });
});