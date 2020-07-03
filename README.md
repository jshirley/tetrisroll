# tetrisroll

A tetris-style randomizer for rolling through sets of keys, to ensure more equitable distribution.

```
  keys := []string{"Brawne", "Sol", "Martin", "Fedmahn", "Het", "Meina"}
  roller := New TetrisRoll(keys)

  // Must iterate through each key before you get a new set!
  key := roller.Roll()
  key := roller.Roll()
  key := roller.Roll()
  key := roller.Roll()
  key := roller.Roll()
  key := roller.Roll()

  // Resets now to the full set
```

## Motivation

Tanya Reilly talks about [Glue Work](https://noidea.dog/glue), and this is the summary:

> Every senior person in an organisation should be aware of the less glamorous - and often less-promotable - work that needs to happen to make a team successful. Managed deliberately, glue work demonstrates and builds strong technical leadership skills. Left unconscious, it can be career limiting. It can push people into less technical roles and even out of the industry.

One thing that bothers me personally is about how it can be career limiting, and often times the way this manifests is
that when a URM performs these tasks is told they should be more technical, or ensure they have more technical impact.

A good way of exposing and tracking this work is through the idea of "Glue Work Distribution". This idea was written up 
at work (Stripe) a while ago by [Amy Nguyen](https://twitter.com/amyngyn) (who is one of those people that we'll all end
up working for someday, she's just that good). The initial setup is to use Slack's slackbot responses and rely on the
random number generator.

Random distribution is not equitable distribution, though. In some settings, it was clear that this distrubtion was
skewed, and significantly so. In one example, one person was selected 8 out of 18 rolls (with 2 people in the set of 5
being selected once and zero!).

[Tetris-style randomization](https://tetris.fandom.com/wiki/Random_Generator) ensures that every item in the set is
picked before the selection restarts. That's why I built this, is to hook it up to a more equitable system to both track
and persist how we can more equitably distribute glue work.

I will likely hook this up to a Slack slash command, and if there is interest will try to make that an open source
library as well!
