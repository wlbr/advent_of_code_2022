



(defun day1 (filename groupsize)
  (let ((x (open filename))
        (partsum 0)
        (max '()))
    (when x
      (loop for line = (read-line x nil)
        while line do 
        (if (string= "" line)
            (progn (if (> groupsize (length max))
                       (setq max (cons partsum max))
                       (let ((pos (position partsum max :test #'(lambda (x y) (> x y)))))
                         (when pos (setf (nth pos max) partsum))))
              (setq max (sort max #'<))
              (setq partsum 0))
            (incf partsum (parse-integer line))))
        (let ((pos (position partsum max :test #'(lambda (x y) (> x y)))))
          (when pos (setf (nth pos max) partsum)))        
      (close x))
    (apply #'+ max)))


(print (day1 "input.txt" 1))
(print (day1 "input.txt" 3))

