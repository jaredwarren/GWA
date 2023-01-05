function testing(...args) {
    console.log("TESTING!!!")
}

function collapsePanel(id) {
    debugger;
    console.log("TODO: colapse panel:", id);
}
function expandPanel(id) {
    debugger;
    console.log("TODO: colapse panel:", id);
}

function toggleNode(el) {
    if ($(el).next('ul').css("visibility") == "hidden") {
        $(el).next('ul').css('visibility', '');
        $(el).next('ul').css("height", 'auto');
        $(el.firstElementChild).removeClass('fa-folder-plus');
        $(el.firstElementChild).addClass('fa-folder-open');
    } else {
        $(el).next('ul').css('visibility', 'hidden');
        $(el).next('ul').css("height", '0px');
        $(el.firstElementChild).removeClass('fa-folder-open');
        $(el.firstElementChild).addClass('fa-folder-plus');
    }
}

function submitForm(id, e) {
    e.preventDefault();
    var f = document.getElementById(id)
    var form = $(f);
    $.ajax({
        type: form.attr('method'),
        url: form.attr('action'),
        data: form.serialize(),
        success: function (data) {
            alert(data);
        }
    });
}