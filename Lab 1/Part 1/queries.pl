% 1 - Simple queries

% Bard's difficulty
?- difficulty('Bard', Diff).

% Sorcerer's primary stat
?- primary_stat('Sorcerer', Stat).

% 2 - Queries with logical operators

% Melee DEX class
?- class(Class), primary_fighting_style(Class, 'melee'), primary_stat(Class, 'DEX').

% All INT or WSD classes
?- findall(Class, (primary_stat(Class, 'INT') ; primary_stat(Class, 'WSD')), Classes).

% 3 - Queries that utilize variables

% Class with difficulty > 5
?- difficulty(Class, Diff), Diff > 5.

% 4 - Queries that utilize rules

% All beginner-friedly classes
?- findall(Class, beginner_friendly(Class), Classes).

% Not beginner-friendly militant class
?- militant(Class), \+ beginner_friendly(Class).
