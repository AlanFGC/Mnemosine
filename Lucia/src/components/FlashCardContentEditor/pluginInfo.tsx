import { QuestionMarkCircleIcon } from '@heroicons/react/20/solid';

export const eventTypes = {
  paragraph: 'paragraph',
  h1: 'h1',
  h2: 'h2',
  ul: 'ul',
  ol: 'ol',
  quote: 'quote',
  formatCode: 'formatCode',
  formatUndo: 'formatUndo',
  formatRedo: 'formatRedo',
  formatBold: 'formatBold',
  formatItalic: 'formatItalic',
  formatUnderline: 'formatUnderline',
  formatStrike: 'formatStrike',
  formatInsertLink: 'formatInsertLink',
  formatAlignLeft: 'formatAlignLeft',
  formatAlignCenter: 'formatAlignCenter',
  formatAlignRight: 'formatAlignRight',
  insertImage: 'insertImage',
  insertToken: 'insertToken',
};

const pluginsList = [
  // Create AnswerToken
  {
    id: 1,
    Icon: QuestionMarkCircleIcon,
    event: eventTypes.insertToken,
  },
];

export default pluginsList;
