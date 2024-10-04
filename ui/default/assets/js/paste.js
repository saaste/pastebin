let copyElement, copiedElement, fadeTimer;

export const initializePaste = () => {
    copyElement = document.getElementById("copy");
    copiedElement = document.getElementById("copied");
    if (copyElement && copiedElement) {
        copyElement.addEventListener("click", copyClicked);
    }
}

const copyClicked = (e) => {
    let rawContent = document.getElementById("raw-content");
    navigator.clipboard.writeText(rawContent.innerText);
    if (!fadeTimer) {
        copiedElement.classList.remove("fadeout");
        copiedElement.classList.add("fadeout");
        fadeTimer = setTimeout(() => {
            copiedElement.classList.remove("fadeout");
            fadeTimer = null;
        }, 3000)
    }
}