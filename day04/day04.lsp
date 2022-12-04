(defun parseline (line) 
  (let* ((hyphen1 (position #\- line))
         (comma (position #\, line))
         (hyphen2 (position #\- line :start comma)))
    (list (list (parse-integer (subseq line 0 hyphen1))
                (parse-integer (subseq line (1+ hyphen1) comma)))
          (list (parse-integer (subseq line (1+ comma) hyphen2))
                (parse-integer (subseq line (1+ hyphen2)))))))


(defun explode (tupel)
  (do* ((i (first tupel) (1+ i))
        (ex (list i) (append ex (list i))))
       ((>= i (second  tupel)) ex))
  )


(defun day4 (filename testfn)
  (let ((x (open filename))
        (score 0))
    (when x
      (loop for line = (read-line x nil)
        while line do 
        (let* ((assignments (parseline line))
               (set1 (explode (first assignments)))
               (set2 (explode (second assignments))))
          (if (funcall testfn set1 set2)
              (incf score))
          )))
    score))


(print (day4 "input.txt" #'(lambda (a b) (or (subsetp a b) (subsetp b a)))))
(print (day4 "input.txt" #'(lambda (a b) (not (null (intersection a b))))))

