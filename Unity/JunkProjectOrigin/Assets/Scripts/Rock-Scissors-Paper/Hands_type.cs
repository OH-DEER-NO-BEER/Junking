using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class Hands_type : MonoBehaviour
{
    public enum Hands
    {
        Rock = 1,
        Scissors,
        Paper
    }

    public static Hands Trans_Hand(int num)
    {
        if (num == 1)
        {
            return Hands.Rock;
        }
        else if (num == 2)
        {
            return Hands.Scissors;
        }
        else
        {
            return Hands.Paper;
        }
    }

    public static string trans_type(Hands hand)
    {
        if (hand == Hands.Rock)
        {
            return "rock";
        }else if (hand == Hands.Scissors)
        {
            return "scissors";
        }else
        {
            return "paper";
        }
    }

    public static int trans_number(string hand_name)
    {
        if (hand_name == "paper")
        {
            return 2;
        }
        else if (hand_name == "scissors")
        {
            return 1;
        }
        else
        {
            return 0;
        }
    }

}
