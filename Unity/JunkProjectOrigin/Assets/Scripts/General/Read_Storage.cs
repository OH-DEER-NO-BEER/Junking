using MiniJSON;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class Read_Storage : MonoBehaviour
{

    private string[] CheckInResponse_func(string data)
    {
        string[] CheckInResponse = new string[5];
        var json = Json.Deserialize(data) as Dictionary<string, object>;
        CheckInResponse[0] = (string)json["message"];

        var mdata = (Dictionary<string, object>)json["myself"];
        CheckInResponse[1] = (string)mdata["name"];
        var mrate = (float[])mdata["rate"];
        for (int i = 0; i < mrate.Length; i++)
        {
            CheckInResponse[2 + i] = mrate[i].ToString();
        }
        //CheckInResponse[5] = mdata["rank"].ToString();

        return CheckInResponse;
    }

    public static string[] DeckInAnnounce_func(string data)
    {
        string[] DeckInAnnounce = new string[5];
        var json = Json.Deserialize(data) as Dictionary<string, object>;
        DeckInAnnounce[0] = (string)json["message"];

        var mdata = (Dictionary<string, object>)json["opponent"];
        DeckInAnnounce[1] = (string)mdata["name"];
        var mrate = (float[])mdata["rate"];
        for (int i = 0; i < mrate.Length; i++)
        {
            DeckInAnnounce[2 + i] = mrate[i].ToString();
        }
        //CheckInResponse[5] = mdata["rank"].ToString();

        return DeckInAnnounce;
    }
}
