import katex from '@traptitech/markdown-it-katex'
import MarkdownIt from 'markdown-it'
import markdownItPrism from 'markdown-it-prism'

import 'prismjs/themes/prism-okaidia.css'
import 'katex/dist/katex.css'

import 'prismjs/components/prism-bash'
import 'prismjs/components/prism-python'
import 'prismjs/components/prism-shell-session'

type Props = {
  content: string
}
export const MarkDown = (props: Props) => {
  const md = new MarkdownIt({
    breaks: true,
  })
  md.use(markdownItPrism)
  md.use(katex, { output: 'html' })

  md.block.State.prototype.skipEmptyLines = function skipEmptyLines(
    from: number
  ): number {
    for (let max = this.lineMax; from < max; from++) {
      if (this.bMarks[from]! + this.tShift[from]! < this.eMarks[from]!) {
        break
      }
      this.push('hardbreak', 'br', 0)
    }
    return from
  }

  const renderedMd = md.render('$f(x)$')

  return <div dangerouslySetInnerHTML={{ __html: renderedMd }} />
}
