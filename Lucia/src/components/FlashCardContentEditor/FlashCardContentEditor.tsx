import { useContext } from 'react';
import { LexicalComposer } from '@lexical/react/LexicalComposer';
import { HistoryPlugin } from '@lexical/react/LexicalHistoryPlugin';
import { PlainTextPlugin } from '@lexical/react/LexicalPlainTextPlugin';
import { ContentEditable } from '@lexical/react/LexicalContentEditable';
import LexicalErrorBoundary from '@lexical/react/LexicalErrorBoundary';
import OnChangePluging from './Plugins/OnChangePluging';
import { FlashCardEditorContext } from '../Contexts/FlashCardEditorContext';

const theme = {};

type FlashCardContentEditorProps = {
  id: string;
};

function onError(error: Error) {
  console.error(error);
}

// eslint-disable-next-line @typescript-eslint/no-unused-vars
function FlashCardContentEditor({ id }:FlashCardContentEditorProps) {
  const { dispatch } = useContext(FlashCardEditorContext);
  const initialConfig = {
    namespace: 'FlashCardContextEditor',
    theme,
    onError,
  };

  function onChange(lexicalState: any) {
    const lexicalStateJSON = lexicalState.toJSON();
    const lexicalStateString = JSON.stringify(lexicalStateJSON);
    dispatch({
      type: 'setEditorContent',
      payload: {
        content: lexicalStateString,
      },
    });
  }

  return (
    <div style={{ backgroundColor: 'gray', height: '400px' }}>
      <LexicalComposer initialConfig={initialConfig}>
        <PlainTextPlugin
          contentEditable={<ContentEditable />}
          placeholder={<div>Enter some text...</div>}
          ErrorBoundary={LexicalErrorBoundary}
        />
        <HistoryPlugin />
        <OnChangePluging onChange={onChange} />
      </LexicalComposer>
    </div>
  );
}

export default FlashCardContentEditor;
