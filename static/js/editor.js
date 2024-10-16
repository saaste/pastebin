let aceEditor, contentElement, syntaxSelectElement, isPublicCheckboxElement, publicPathContainer, deleteButton, wrapCheckBoxElement, documentForm, nameInput;

export const initializeEditor = () => {
    const editorElement = document.getElementById("editor");
    contentElement = document.getElementById("content");
    syntaxSelectElement = document.getElementById("syntax");
    isPublicCheckboxElement = document.getElementById("is_public");
    publicPathContainer = document.querySelector("div.public-path-container");
    deleteButton = document.getElementById("delete");
    wrapCheckBoxElement = document.getElementById("wrap");
    documentForm = document.getElementById("document-form");
    nameInput = document.getElementById("name");

    if (!editorElement || !contentElement || !syntaxSelectElement | !isPublicCheckboxElement | !publicPathContainer | !wrapCheckBoxElement) {
        return;
    }

    let wrapEnabled = localStorage.getItem("wrap") == "true" ? true : false;
    wrapCheckBoxElement.checked = wrapEnabled;

    let selectedSyntax = syntaxSelectElement.value;
    syntaxSelectElement.addEventListener("change", syntaxChanged);

    aceEditor = ace.edit("editor");
    aceEditor.setOptions({
        printMargin: false,
        fontFamily: "Roboto Mono",
        fontSize: "1rem",
        useWorker: false,
        wrap: wrapEnabled,
    });
    aceEditor.setTheme("ace/theme/monokai");
    aceEditor.session.setMode("ace/mode/" + selectedSyntax);
    aceEditor.session.on("change", editorChanged);

    isPublicCheckboxElement.addEventListener("change", isPublicChanged)
    if (!isPublicCheckboxElement.checked) {
        publicPathContainer.classList.add("hidden");
        aceEditor.resize();
    }

    wrapCheckBoxElement.addEventListener("change", wrapCheckBoxElementChanged)

    if (deleteButton) {
        deleteButton.addEventListener("click", deleteButtonClicked);
    }

    if (documentForm && nameInput) {
        nameInput.addEventListener("change", nameInputChanged)
        nameInput.addEventListener("keyup", nameInputChanged)
    }

}

const editorChanged = (delta) => {
    contentElement.value = aceEditor.getValue();
}

const syntaxChanged = (e) => {
    aceEditor.session.setMode("ace/mode/" + e.target.value);
}

const isPublicChanged = (e) => {
    if (e.target.checked) {
        publicPathContainer.classList.remove("hidden");
    } else {
        publicPathContainer.classList.add("hidden");
    }
    aceEditor.resize();
}

const deleteButtonClicked = (e) => {
    let result = confirm("Are you sure you want to delete this?")
    if (!result) {
        e.preventDefault();
    }
}

const wrapCheckBoxElementChanged = (e) => {
    aceEditor.setOption("wrap", e.target.checked)
    localStorage.setItem("wrap", e.target.checked)
}

const nameInputChanged = (e) => {
    let title = document.title;
    let parts = title.split(" | ");
    let siteName = parts[parts.length - 1]

    if (e.target.value.length > 0) {
        document.title = `${e.target.value} | ${siteName}`
    } else {
        document.title = siteName;
    }
    
}