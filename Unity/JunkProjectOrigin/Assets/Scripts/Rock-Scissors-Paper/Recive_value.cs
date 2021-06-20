using MiniJSON;
using System;
using System.Collections;
using System.Collections.Generic;
using System.Runtime.InteropServices;
using UnityEngine;
using UnityEngine.SceneManagement;
using UnityEngine.UI;
using static Hands_type;

public class Recive_value : MonoBehaviour
{
    [SerializeField]
    GameObject Player1;
    [SerializeField]
    GameObject Player2;

    [SerializeField]
    GameObject Select_Time;
    [SerializeField]
    GameObject Timer;
    [SerializeField]
    List<GameObject> objs;

    [SerializeField]
    GameObject Status;

    [SerializeField]
    GameObject Error_text;

    AudioSource audioSource;

    private string json_text="";
    private string[] Responce_data;

    [HideInInspector]
    public string player_status;

    //JS -> SendMessage("EventControl","SetJson",string data)
    public void SetJson(string text)
    {
        json_text = text;
    }

    [DllImport("__Internal")]
    private static extern void SetLocalStorage(string key, string value);

    private void Send_Select_Hands(Hands hand)
    {
#if !UNITY_EDITOR
        SetLocalStorage("selectHand", trans_type(hand));
#endif
        SceneManager.LoadScene("Hands_Animation_Scene");
    }

    private void GameSceneLoaded(Scene next, LoadSceneMode mode)
    {
        var gameManager = GameObject.Find("EventControl").GetComponent<Animation_Control>();
        gameManager.player_status = player_status;
        SceneManager.sceneLoaded -= GameSceneLoaded;
    }


    [HideInInspector]
    public string self_data = null;

    private string[] DeckInAnnounce_func(string data)
    {
        string[] DeckInAnnounce = new string[10];
        var json = Json.Deserialize(data) as Dictionary<string, object>;
        DeckInAnnounce[0] = (string)json["message"];
        Debug.Log(DeckInAnnounce[0]);

        var mdata = (Dictionary<string, object>)json["p1"];
        DeckInAnnounce[1] = (string)mdata["name"];
        var mrate = (Dictionary<string, object>)mdata["rate"];
        DeckInAnnounce[2] = mrate["rock"].ToString();
        DeckInAnnounce[3] = mrate["scissors"].ToString();
        DeckInAnnounce[4] = mrate["paper"].ToString();
        //CheckInResponse[5] = mdata["rank"].ToString();

        mdata = (Dictionary<string, object>)json["p2"];
        DeckInAnnounce[5] = (string)mdata["name"];
        mrate = (Dictionary<string, object>)mdata["rate"];
        DeckInAnnounce[6] = mrate["rock"].ToString();
        DeckInAnnounce[7] = mrate["scissors"].ToString();
        DeckInAnnounce[8] = mrate["paper"].ToString();
        //DeckInAnnounce[5] = mdata["rank"].ToString();

        return DeckInAnnounce;
    }

    private void Shader_change()
    {
        Color color = objs[0].transform.Find("r_handMeshNode").GetComponent<Renderer>().material.color;
        color = new Color(180, 88, 88);
        color.a = 155;

        color = objs[1].transform.Find("r_handMeshNode").GetComponent<Renderer>().material.color;
        color = new Color(50, 48, 144);
        color.a = 155;

        color = objs[1].transform.Find("r_handMeshNode").GetComponent<Renderer>().material.color;
        color = new Color(173, 172, 110);
        color.a = 155;

        color = Player1.transform.Find("P_Back").GetComponent<Image>().color;
        color = new Color(255, 255, 255);
        color.a = 100;

        color = Player2.transform.Find("P_Back").GetComponent<Image>().color;
        color = new Color(255, 255, 255);
        color.a = 100;
    }

    private void Recieve_data(GameObject player1,GameObject player2, string[] data)
    {
        if (player_status == "Room Made")
        {
            player1.GetComponent<Text>().text = data[1] + Environment.NewLine +
                                                "Rock:" + data[2].ToString() + Environment.NewLine +
                                                "Scissors:" + data[3].ToString() + Environment.NewLine +
                                                "Paper:" + data[4].ToString() + Environment.NewLine;
            player2.GetComponent<Text>().text = data[5] + Environment.NewLine +
                                               "Rock:" + data[6].ToString() + Environment.NewLine +
                                               "Scissors:" + data[7].ToString() + Environment.NewLine +
                                               "Paper:" + data[8].ToString() + Environment.NewLine;
        }else if(player_status == "Room Entered")
        {
            player1.GetComponent<Text>().text = data[5] + Environment.NewLine +
                                    "Rock:" + data[6].ToString() + Environment.NewLine +
                                    "Scissors:" + data[7].ToString() + Environment.NewLine +
                                    "Paper:" + data[8].ToString() + Environment.NewLine;
            player2.GetComponent<Text>().text = data[1] + Environment.NewLine +
                                               "Rock:" + data[2].ToString() + Environment.NewLine +
                                               "Scissors:" + data[3].ToString() + Environment.NewLine +
                                               "Paper:" + data[4].ToString() + Environment.NewLine;
        }
        else
        {
            player1.GetComponent<Text>().text = "Load Error";
            player2.GetComponent<Text>().text = "Load Error";
        }
    }

    private void Valid_obj()
    {
        Timer.GetComponent<TimerCtrl>().enabled = true;
        foreach (GameObject obj in objs)
        {
            obj.GetComponent<Motion_and_Event>().enabled = true;
            obj.GetComponent<Motion_and_Event>().mouse_flg = true;
        }
        Select_Time.SetActive(true);
    }

    // Start is called before the first frame update
    void Start()
    {

        Timer.GetComponent<TimerCtrl>().enabled = false;
        foreach(GameObject obj in objs)
        {
            obj.GetComponent<Motion_and_Event>().enabled = false;
            obj.GetComponent<Motion_and_Event>().mouse_flg = false;
        }
        Select_Time.SetActive(false);

        audioSource = gameObject.GetComponent<AudioSource>();
        SceneManager.sceneLoaded += GameSceneLoaded;

#if UNITY_EDITOR
        json_text = (Resources.Load("new") as TextAsset).text;
#endif
    }

    // Update is called once per frame
    void Update()
    {

        if (Timer.GetComponent<TimerCtrl>().selected)
        {
            Send_Select_Hands(Timer.GetComponent<TimerCtrl>().hand);
        }

        foreach (GameObject obj in objs)
        {
            if (obj.GetComponent<Motion_and_Event>().selected)
            {
                Send_Select_Hands(obj.GetComponent<Motion_and_Event>().hand);
            }
        }

        if (json_text == "")
        {
            return;
        }

        Responce_data = DeckInAnnounce_func(json_text);

        player_status = Responce_data[0];

        if (player_status == "Room Occupied")
        {
            Error_text.GetComponent<Text>().text = "An error has occurred";
            json_text = "";
        }
        else if(player_status == "Room Made" || player_status == "Room Entered")
        {
            Recieve_data(Player1,Player2, Responce_data);
            Shader_change();
            Valid_obj();

            audioSource.Play();
            json_text = "";
        }
        else
        {
            Error_text.GetComponent<Text>().text = "An unexplained error has occurred.";
            json_text = "";
        }

    }
}
