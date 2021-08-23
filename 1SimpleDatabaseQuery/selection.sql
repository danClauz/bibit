SELECT us.ID, us.UserName, p.ParentUserName
FROM User us
    LEFT JOIN Parent p ON p.ID = us.ID