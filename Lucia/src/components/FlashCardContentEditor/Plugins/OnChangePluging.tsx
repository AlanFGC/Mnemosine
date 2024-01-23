/* eslint-disable @typescript-eslint/no-unused-vars */
import { useLexicalComposerContext } from '@lexical/react/LexicalComposerContext';
import { useEffect } from 'react';

function onChangePluging({ onChange }: { onChange: (value: unknown) => void }) {
  // eslint-disable-next-line react-hooks/rules-of-hooks
  const [editor] = useLexicalComposerContext();
  // eslint-disable-next-line react-hooks/rules-of-hooks
  useEffect(() => editor.registerUpdateListener(({ editorState }) => {
    onChange(editorState);
  }), [editor, onChange]);
}

export default onChangePluging;
