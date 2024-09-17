% 1 - Single-argument facts

% i - Character classes
class('Bard').
class('Barbarian').
class('Fighter').
class('Wizard').
class('Druid').
class('Monk').
class('Rogue').
class('Paladin').
class('Ranger').
class('Artificer').
class('Cleric').
class('Warlock').
class('Sorcerer').
class('Alchemist').
class('Warlord').
class('Jaeger').
class('Stargazer').
class('Blood Hunter').
class('Runekeeper').
class('Shaman').

% 2 - Facts with 2 arguments

% i - Primary stats for classes
primary_stat('Bard', 'CHR').
primary_stat('Barbarian', 'STR').
primary_stat('Fighter', 'DEX').
primary_stat('Wizard', 'INT').
primary_stat('Druid', 'WSD').
primary_stat('Cleric', 'WSD').
primary_stat('Artificer', 'INT').
primary_stat('Warlock', 'CHR').
primary_stat('Monk', 'DEX').
primary_stat('Paladin', 'STR').
primary_stat('Rogue', 'DEX').
primary_stat('Ranger', 'DEX').
primary_stat('Sorcerer', 'CHR').
primary_stat('Alchemist', 'INT').
primary_stat('Warlord', 'STR').
primary_stat('Jaeger', 'CNS').
primary_stat('Stargazer', 'INT').
primary_stat('Blood Hunter', 'DEX').
primary_stat('Runekeeper 'INT').
primary_stat('Shaman 'WSD').

% ii - Primary fighting styles for classes
primary_fighting_style('Bard', 'support').
primary_fighting_style('Barbarian', 'melee').
primary_fighting_style('Fighter', 'ranged').
primary_fighting_style('Wizard', 'caster').
primary_fighting_style('Druid', 'caster').
primary_fighting_style('Monk', 'melee').
primary_fighting_style('Rogue', 'melee').
primary_fighting_style('Paladin', 'swordlemage').
primary_fighting_style('Ranger', 'ranged').
primary_fighting_style('Artificer', 'caster').
primary_fighting_style('Cleric', 'support').
primary_fighting_style('Warlock', 'caster').
primary_fighting_style('Sorcerer', 'caster').
primary_fighting_style('Alchemist', 'support').
primary_fighting_style('Warlord', 'support').
primary_fighting_style('Jaeger', 'melee').
primary_fighting_style('Stargazer', 'caster').
primary_fighting_style('Blood Hunter', 'swordlemage').
primary_fighting_style('Runekeeper 'caster').
primary_fighting_style('Shaman 'caster').

% iii - Difficulty scale for classes
difficulty('Bard', 7).
difficulty('Barbarian', 2).
difficulty('Fighter', 1).
difficulty('Wizard', 9).
difficulty('Druid', 7).
difficulty('Monk', 4).
difficulty('Rogue', 2).
difficulty('Paladin', 3).
difficulty('Ranger', 3).
difficulty('Artificer', 6).
difficulty('Cleric', 5).
difficulty('Warlock', 6).
difficulty('Sorcerer', 10).
difficulty('Alchemist', 8).
difficulty('Warlord', 2).
difficulty('Jaeger', 2).
difficulty('Stargazer', 9).
difficulty('Blood Hunter', 6).
difficulty('Runekeeper', 4).
difficulty('Shaman', 8).

% 3 - Rules

% Class is beginner-friendly if it has difficulty <= 3
beginner_friendly(Class) :-
	difficulty(Class, Difficulty),
	Difficulty =< 3.

% Class challenging if it has difficulty >= 8
challenging(Class) :-
	difficulty(Class, Difficulty),
	Difficulty >= 8.
	
% Militant classes are those that fight in melee or with ranged weapons (swordlemage fits here)
militant(Class) :-
	primary_fighting_style(Class, 'melee');
	primary_fighting_style(Class, 'ranged');
	primary_fighting_style(Class, 'swordlemage').
	
% Magic classes are those that cast spells or support their partymates (swordlemage fits here as well)
magic(Class) :-
	primary_fighting_style(Class, 'caster');
	primary_fighting_style(Class, 'support');
	primary_fighting_style(Class, 'swordlemage').
	
% Body classes are those whose primary stat is either STR, DEX or CNS
body(Class) :-
	primary_stat(Class, 'STR');
	primary_stat(Class, 'DEX');
	primary_stat(Class, 'CNS').
	
% Soul classes are those whose primary stat is either INT, WSD or CHR
soul(Class) :-
	primary_stat(Class, 'INT');
	primary_stat(Class, 'WSD');
	primary_stat(Class, 'CHR').