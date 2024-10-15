let aceRawContent, rawContentElement;

export const initializeRawContent = () => {
    rawContentElement = document.getElementById("raw-content");
    if (!rawContentElement) {
        return;
    }

    let syntax = rawContentElement.dataset.syntax || "text"

    aceRawContent = ace.edit("raw-content");
    aceRawContent.setOptions({
        printMargin: false,
        fontFamily: "Roboto Mono",
        fontSize: "1rem",
        useWorker: false,
        wrap: true,
        readOnly: true,
        showGutter: false,
    });
    aceRawContent.setTheme("ace/theme/monokai");
    aceRawContent.session.setMode("ace/mode/" + syntax);
}