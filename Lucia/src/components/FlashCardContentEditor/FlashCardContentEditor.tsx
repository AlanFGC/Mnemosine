import { useEffect } from 'react';

import { LexicalComposer } from '@lexical/react/LexicalComposer';
import { RichTextPlugin } from '@lexical/react/LexicalRichTextPlugin';
import { ContentEditable } from '@lexical/react/LexicalContentEditable';
import { HistoryPlugin } from '@lexical/react/LexicalHistoryPlugin';
import { useLexicalComposerContext } from '@lexical/react/LexicalComposerContext';
import LexicalErrorBoundary from '@lexical/react/LexicalErrorBoundary';

function MyCustomAutoFocusPlugin() {
  const [editor] = useLexicalComposerContext();

  useEffect(() => {
    // Focus the editor when the effect fires!
    editor.focus();
  }, [editor]);

  return null;
}

function onError(error: Error) {
  // eslint-disable-next-line no-console
  console.error(error);
}

interface FlashCardContentEditorProps {
  setCardContent: (content: string) => void;
  addAnswerToken: (tokens: number) => void;
  removeAnswerToken: (tokens: number) => void;
}

// TODO add the logic to the tokens

function FlashCardContentEditor(
  {
    setCardContent,
    addAnswerToken,
    removeAnswerToken,
  }: FlashCardContentEditorProps,
) {
  const initialConfig = {
    namespace: 'MyEditor',
    onError,
  };

  return (
    <LexicalComposer initialConfig={initialConfig}>
      <RichTextPlugin
        contentEditable={<ContentEditable />}
        placeholder={<h3>Enter some text...</h3>}
        ErrorBoundary={LexicalErrorBoundary}
      />
      <HistoryPlugin />
      <MyCustomAutoFocusPlugin />
    </LexicalComposer>
  );
}

export default FlashCardContentEditor;
