let aceEditor, contentElement, syntaxSelectElement, isPublicCheckboxElement, publicPathContainer, deleteButton, copyElement, copiedElement, fadeTimer;

// window.onload = () => {
//     initializePaste();
//     initializeEditor();
// }

// const initializePaste = () => {
//     copyElement = document.getElementById("copy");
//     copiedElement = document.getElementById("copied");
//     if (copyElement && copiedElement) {
//         copyElement.addEventListener("click", copyClicked);
//     }
// }

export const initializeEditor = () => {
    const editorElement = document.getElementById("editor");
    contentElement = document.getElementById("content");
    syntaxSelectElement = document.getElementById("syntax");
    isPublicCheckboxElement = document.getElementById("is_public");
    publicPathContainer = document.querySelector("div:has(#public_path)")
    deleteButton = document.getElementById("delete")

    if (!editorElement || !contentElement || !syntaxSelectElement | !isPublicCheckboxElement | !publicPathContainer) {
        console.log("Editor not defined. Skipping...")
        return;
    }

    let selectedSyntax = syntaxSelectElement.value;
    syntaxSelectElement.addEventListener("change", syntaxChanged);

    aceEditor = ace.edit("editor");
    aceEditor.setOptions({
        printMargin: false,
        fontFamily: "Roboto Mono",
        fontSize: "1rem",
        useWorker: false,
    });
    aceEditor.setTheme("ace/theme/monokai");
    aceEditor.session.setMode("ace/mode/" + selectedSyntax);
    aceEditor.session.on("change", editorChanged);

    isPublicCheckboxElement.addEventListener("change", isPublicChanged)
    if (!isPublicCheckboxElement.checked) {
        publicPathContainer.classList.add("hidden");
    }

    if (deleteButton) {
        deleteButton.addEventListener("click", deleteButtonClicked);
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
}

const deleteButtonClicked = (e) => {
    let result = confirm("Are you sure you want to delete this?")
    if (!result) {
        e.preventDefault();
    }
}

// const copyClicked = (e) => {
//     let rawContent = document.getElementById("raw-content");
//     navigator.clipboard.writeText(rawContent.innerText);
//     if (!fadeTimer) {
//         copiedElement.classList.remove("fadeout");
//         copiedElement.classList.add("fadeout");
//         fadeTimer = setTimeout(() => {
//             copiedElement.classList.remove("fadeout");
//             fadeTimer = null;
//         }, 3000)
//     }
// }