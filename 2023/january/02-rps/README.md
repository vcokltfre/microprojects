# January 01, 2023 - Rock, Paper, Scissors

Another simple and classic project for this one: Rock, Paper, Scissors. The rules of the game are simpe, two players choose either rock, paper, or scissors, and the winner is decided by the following rules:

- Rock beats scissors
- Scissors beats paper
- Paper beats rock

As with Fizzbuzz, I'm not content with the most basic possible implementation, so I'm going to make it extensible with a custom list of objects. The list will be ordered as the item with index `n` beating the item with index `n-1`, except for the first item, which beats the last. For example, if the list is `["Rock", "Paper", "Scissors", "Lizard", "Spock"]`, then the rules are:

- Rock beats spock
- Spock beats lizard
- Lizard beats scissors
- Scissors beats paper
- Paper beats rock
- All other combinations are a draw

This will be implemented as a single-player game using pseudo-random numbers to emutale a second player.
