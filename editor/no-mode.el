;;; no-mode.el --- Major Mode for editing Nolang source code -*- lexical-binding: t -*-

;; Copyright (c) 2022 Roberto Hermenegildo Dias

;; Author: Roberto Hermenegildo Dias
;; URL: https://github.com/robertosixty1/nolang

;; Permission is hereby granted, free of charge, to any person obtaining a copy
;; of this software and associated documentation files (the "Software"), to deal
;; in the Software without restriction, including without limitation the rights
;; to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
;; copies of the Software, and to permit persons to whom the Software is
;; furnished to do so, subject to the following conditions:

;; The above copyright notice and this permission notice shall be included in all
;; copies or substantial portions of the Software.

;; THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
;; IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
;; FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
;; AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
;; LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
;; OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
;; SOFTWARE.

;;; Commentary:
;;
;; Major Mode for editing Nolang source code.

(defconst no-mode-syntax-table
  (with-syntax-table (copy-syntax-table)
    ;; C/C++ style comments
    (modify-syntax-entry ?/ ". 124b")
    (modify-syntax-entry ?* ". 23")
    (modify-syntax-entry ?\n "> b")
    ;; Chars are the same as strings
    (modify-syntax-entry ?' "\"")
    (syntax-table))
  "Syntax table for `no-mode'.")

(eval-and-compile
  (defconst no-keywords
    '("if" "else" "end" "while" "do" "macro" "include" "const"
      "increment" "reset" "memory" "here" "let" "in" "func" "done")))

(defconst no-highlights
  `((,(regexp-opt no-keywords 'symbols) . font-lock-keyword-face)))

;;;###autoload
(define-derived-mode no-mode prog-mode "no"
  "Major Mode for editing Nolang source code"
  :syntax-table no-mode-syntax-table
  (setq font-lock-defaults '(no-highlights))
  (setq-local comment-start "// "))

;;;###autoload
(add-to-list 'auto-mode-alist '("\\.no\\'" . no-mode))

(provide 'no-mode)

;;; no-mode.el ends here
