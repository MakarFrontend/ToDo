window.onload = async function() {
    let res = await fetch("/content");
    let rt = await res.text();
    document.getElementById("main").innerHTML = rt;
};


async function getUser(login, pass) {
    if (login != "" && pass != "") {
        let res = await fetch("/user/" + login + "?" + "password=" + pass)
        if (res.status === 200) {
            document.getElementById("main").innerHTML = await res.text();
            sessionStorage.setItem("login", login);
            sessionStorage.setItem("pass", pass);
        } else {
            let resTXT = await res.text()
            if (res.status === 404) {
                alert("Неправильный логин или пароль!")
            } else {
                alert(resTXT)
            }
        }
    }
}

async function exit() {
    sessionStorage.removeItem("login");
    sessionStorage.removeItem("pass");

    let res = await fetch("/content");
    let rt = await res.text();
    document.getElementById("main").innerHTML = rt;
}

async function start() {
    if (sessionStorage.getItem("login") == null && sessionStorage.getItem("pass") == null) {
        let login = prompt("Логин");
        let password = prompt("Пароль");
        await getUser(login, password);
    } else {
        await getUser(sessionStorage.getItem("login"), sessionStorage.getItem("pass"));
    }
}

async function toggleStatus(id) {
    const log = sessionStorage.getItem("login");
    let resp = await fetch(`/user/${log}/${id}/status`, {method: "POST"});
    let resTXT = await resp.text();
    alert(resTXT);
    if (resp.status == 200) {
        getUser(sessionStorage.getItem("login"), sessionStorage.getItem("pass"));
    }
}

async function registUser() {
    let rex = /^[a-zA-Z0-9]+$/

    const login = document.getElementById("login").value;
    const pass1 = document.getElementById("pass1").value;
    const pass2 = document.getElementById("pass2").value;
    if ((login != "" && pass1 != "" && pass2 != "") && (pass1 == pass2) && (rex.test(login)) && (rex.test(pass1))) {
        let resp = await fetch(`/new/user?password=${pass1}&login=${login}`, {method: "POST"})
        let rTXT = await resp.text()
        alert(rTXT)
    } else {
        alert("Поля в форме регистрации не должны быть пустыми!\nПароли должны совпадать!\nПароль и логин могут содержать только латинские буквы и цифры!")
    }

    document.getElementById("login").value = "";
    document.getElementById("pass2").value = "";
    document.getElementById("pass1").value = "";
}

async function newToDo() {
    let toSend = {};
    toSend.txt = document.getElementById("aboutToDo").value;
    toSend.tag = document.forms['tag']['selectTag'].value;
    toSend.status = false;

    if (toSend.txt === "") {
        alert("Описание задачи не должно быть пустым")
    }

    let strToSend = JSON.stringify(toSend);

    let resp = await fetch(`/user/${sessionStorage.getItem("login")}/new/task`, {
        method: "POST", 
        body: strToSend
    });

    const respText = await resp.text()

    alert(respText)

    if (resp.status = 200) {
        await start()
    }
}

async function deleteAccount() {
    let answ = prompt("Для подтверждения удаления ответьте на вопрос!\n2 + 2 = ");
    if (answ === "4") {
        const login = sessionStorage.getItem("login");
        let resp = await fetch(`/user/${login}/del`, {
            method: "DELETE"
        });

        let rText = await resp.text()

        alert(rText)
        if (resp.status == 200) {
            sessionStorage.clear()
            let res = await fetch("/content");
            let rt = await res.text();
            document.getElementById("main").innerHTML = rt;
        }
    }
}

async function deleteTask(id) {
    let resp = await fetch(`/task/${id}/del`, {
        method: "DELETE"
    });

    let rText = await resp.text()

    alert(rText)
    if (resp.status == 200) {
        await start()
    }
}