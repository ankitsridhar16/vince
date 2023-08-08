import { Site } from "../../../../vince"
import * as monaco from "monaco-editor"
import { CompletionItemKind } from "./types"

export const createSchemaCompletionProvider = (questDBTables: Site[] = []) => {
    const completionProvider: monaco.languages.CompletionItemProvider = {
        triggerCharacters: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz \"".split(
            "",
        ),
        provideCompletionItems(model, position) {
            const word = model.getWordUntilPosition(position)

            const textUntilPosition = model.getValueInRange({
                startLineNumber: 1,
                startColumn: 1,
                endLineNumber: position.lineNumber,
                endColumn: word.startColumn,
            })

            const range = {
                startLineNumber: position.lineNumber,
                endLineNumber: position.lineNumber,
                startColumn: word.startColumn,
                endColumn: word.endColumn,
            }

            const nextChar = model.getValueInRange({
                startLineNumber: position.lineNumber,
                startColumn: word.endColumn,
                endLineNumber: position.lineNumber,
                endColumn: word.endColumn + 1,
            })

            if (
                word.word ||
                /(FROM|INTO|TABLE) $/gim.test(textUntilPosition) ||
                (/'$/gim.test(textUntilPosition) && !textUntilPosition.endsWith("= '"))
            ) {
                const openQuote = textUntilPosition.substr(-1) === "\"";
                const nextCharQuote = nextChar == "\"";
                return {
                    suggestions: questDBTables.map((item) => {
                        return {
                            label: item.domain,
                            kind: CompletionItemKind.Class,
                            insertText:
                                openQuote
                                    ? item.domain + (nextCharQuote ? "" : "\"")
                                    : /^[a-z0-9_]+$/i.test(item.domain) ? item.domain : `"${item.domain}"`,
                            range,
                        }
                    }),
                }
            }
        },
    }

    return completionProvider
}