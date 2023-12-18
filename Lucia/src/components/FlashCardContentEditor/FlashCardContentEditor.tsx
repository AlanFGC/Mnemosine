import { LexicalComposer } from '@lexical/react/LexicalComposer';
import { RichTextPlugin } from '@lexical/react/LexicalRichTextPlugin';
import { ContentEditable } from '@lexical/react/LexicalContentEditable';
import { HistoryPlugin } from '@lexical/react/LexicalHistoryPlugin';
import { useLexicalComposerContext } from '@lexical/react/LexicalComposerContext';
import LexicalErrorBoundary from '@lexical/react/LexicalErrorBoundary';
import { useEffect } from 'react';
import Toolbar from './ToolBar/Toolbar';

const theme = {};
// Catch any errors that occur during Lexical updates and log them
// or throw them as needed. If you don't throw them, Lexical will
// try to recover gracefully without losing user data.
function onError(error: Error) {
  console.error(error);
}

function FlashCardContentEditor() {
  // const [editorState, setEditorState] = useState;

  const initialConfig = {
    namespace: 'MyEditor',
    theme,
    onError,
  };

  function consoleLoggerPlugin(): void {
    const [editor] = useLexicalComposerContext();

    useEffect(() => {
      const removeTextContentListener = editor.registerTextContentListener(
        (textContent: string) => {
          // The latest text content of the editor!
          console.log(textContent);
        },
      );
      return () => {
        removeTextContentListener();
      };
    }, [editor]);
  }

  return (
    <div className="FlashCardConentEditor">
      <Toolbar />
      <LexicalComposer initialConfig={initialConfig}>
        <RichTextPlugin
          contentEditable={<ContentEditable />}
          placeholder={<div>Enter some text...</div>}
          ErrorBoundary={LexicalErrorBoundary}
        />
        <HistoryPlugin />
        <consoleLoggerPlugin />
      </LexicalComposer>
    </div>
  );
}

export default FlashCardContentEditor;
