SELECT us.ID, us.UserName, p.ParentUserName from User us
LEFT JOIN Parent p ON p.ID = us.ID