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

% ii - Stats
stat('STR').
stat('DEX').
stat('CNS').
stat('INT').
stat('WSD').
stat('CHR').

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

% ii - Primary fighting styles for classes
primary_fighting_style('Bard', 'support').
primary_fighting_style('Barbarian', 'melee').
primary_fighting_style('Fighter', 'ranged').
primary_fighting_style('Wizard', 'caster').
primary_fighting_style('Druid', 'caster').
primary_fighting_style('Monk', 'melee').
primary_fighting_style('Rogue', 'melee').
primary_fighting_style('Paladin', 'melee').
primary_fighting_style('Ranger', 'ranged').
primary_fighting_style('Artificer', 'caster').
primary_fighting_style('Cleric', 'support').
primary_fighting_style('Warlock', 'caster').
primary_fighting_style('Sorcerer', 'caster').
primary_stat('Alchemist', 'support').

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
