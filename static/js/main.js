let attention = Prompt();
function notify(msg, msgType){
    notie.alert({
        type: msgType, // optional, default = 4, enum: [1, 2, 3, 4, 5, 'success', 'warning', 'error', 'info', 'neutral']
        text: msg,
        stay: fasle, // optional, default = false
        time: 3, // optional, default = 3, minimum = 1,
        position: top // optional, default = 'top', enum: ['top', 'bottom']
      })
}

function notifyModal(title, text, icon, confirmationButtonText) {
    Swal.fire({
        title: title,
        html: text,
        icon: icon,
        confirmButtonText: confirmButtonText
    })
};

function Prompt() {
    let toast = (c)=> {
        const {
            msg = "",
            icon = "success",
            position="top-end",
        } = c;
        const Toast = Swal.mixin({
            toast: true,
            title: msg,
            position: position,
            showConfirmButton: false,
            timer: 3000,
            timerProgressBar: true,
            didOpen: (toast) => {
              toast.addEventListener('mouseenter', Swal.stopTimer)
              toast.addEventListener('mouseleave', Swal.resumeTimer)
            }
          })
          
          Toast.fire({})
    }
    let success = (c)=> {
        const {
            msg = "",
            title = "",
            footer = "",
        } = c;
        Swal.fire({
            icon: 'success',
            title: title,
            text: msg,
            footer: footer
          })
    }
    let error = (c)=> {
        const {
            msg = "",
            title = "",
            footer = "",
        } = c;
        Swal.fire({
            icon: 'error',
            title: title,
            text: msg,
            footer: footer
          })
    }

    let custom = async (c)=> {
        const {
            msg="",
            title= "",
        } = c;
        const { value: result } = await Swal.fire({
            title: title,
            html: msg,
            backdrop: false,
            showCancelButton: true,
            focusConfirm: false,
            willOpen: () => {
                if (c.willOpen !== undefined) {
                    c.willOpen();
                }
            },
            preConfirm: () => {
              return [
                document.getElementById('start').value,
                document.getElementById('end').value
              ]
            },
            didOpen: () => {
                if (c.didOpen !== undefined) {
                    c.didOpen();
                }
            }
          })
        // if (
        //   c.callback &&
        //   result &&
        //   result.dismiss !== Swal.DismissReason.cacel &&
        //   result.value !== ""
        // ) {
        //   console.log("callback")
        //   c.callback(result);
        // } else {
        //   c.callback(false);
        // }
        // console.log(c.callback)
        if (result) {
            if (result.dismiss !== Swal.DismissReason.cancel) {
                if (result.value !== "") {
                    if (c.callback !== undefined) {
                        //console.log(result)
                        c.callback(result)
                    } else {
                        c.callback(false)
                    }
                } else {
                    c.callback(false)
                }
            }
        }
    }
    return {
        toast: toast,
        success: success,
        error: error,
        custom: custom,
    }
}