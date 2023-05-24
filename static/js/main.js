let attention = Prompt();


// client side validation 
(() => {
    'use strict'
  
    // Fetch all the forms we want to apply custom Bootstrap validation styles to
    const forms = document.querySelectorAll('.needs-validation')
  
    // Loop over them and prevent submission
    Array.from(forms).forEach(form => {
      form.addEventListener('submit', event => {
        if (!form.checkValidity()) {
          event.preventDefault()
          event.stopPropagation()
        }
  
        form.classList.add('was-validated')
      }, false)
    })
})();

// end client side validation

function notify(msg, msgType){
    notie.alert({
        type: msgType,
        text: msg,
        stay: false,
        time: 3,
        position: 'top'
    })
};



function notifyModal(title, text, icon, confirmationButtonText) {
    Swal.fire({
        title: title,
        html: text,
        icon: icon,
        confirmButtonText: confirmButtonText
    })
};

if (document.querySelector('meta[name="message"]')) {
    message = document.querySelector('meta[name="message"]').content
    type = document.querySelector('meta[name="type"]').content
    notify(message, type)
}


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