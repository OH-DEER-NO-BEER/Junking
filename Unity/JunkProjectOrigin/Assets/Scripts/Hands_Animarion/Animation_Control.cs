using MiniJSON;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.SceneManagement;
using UnityEngine.UI;

public class Animation_Control : MonoBehaviour
{
    [SerializeField]
    GameObject subtitle;
    Text text;

    [HideInInspector]
    public string player_status;

    float time;

    int counter = 0;

    [SerializeField]
    GameObject myHand;

    [SerializeField]
    GameObject enemyHand;

    Animator myanimator;
    Animator enemyanimator;

    AudioSource audioSource;
    [SerializeField]
    AudioClip junksound;

    private string json_text = "";
    private string[] Responce_data = null;

    private void GameSceneLoaded(Scene next, LoadSceneMode mode)
    {
        if (next.name == "Junk_Result") {
            var gameManager = GameObject.Find("EventControl").GetComponent<Result_app>();
            gameManager.self_h = Responce_data[0];
            gameManager.oppenent_h = Responce_data[1];
            gameManager.result_Junk = Responce_data[2];
            SceneManager.sceneLoaded -= GameSceneLoaded;
        }
        else
        {
            var gameManager = GameObject.Find("EventControl").GetComponent<Recive_value>();
            gameManager.player_status = player_status;
            SceneManager.sceneLoaded -= GameSceneLoaded;
        }
    }


    //JS -> SendMessage("EventControl","SetJson",string data)
    public void SetJson(string text)
    {
        json_text = text;
    }

    private string[] DeckInAnnounce_func(string data)
    {
        string[] Result_Announce = new string[3];
        var json = Json.Deserialize(data) as Dictionary<string, object>;
        if (player_status == "Room Made")
        {
            Result_Announce[0] = (string)json["P1Hand"];
            Result_Announce[1] = (string)json["P2Hand"];
            Result_Announce[2] = (string)json["result"];
        }else if(player_status == "Room Entered")
        {
            Result_Announce[1] = (string)json["P1Hand"];
            Result_Announce[0] = (string)json["P2Hand"];
            Result_Announce[2] = (string)json["result"];
        }
        else
        {

            Result_Announce[0] = null;
            Result_Announce[1] = null;
            Result_Announce[2] = "Error";
        }
        Debug.Log(Result_Announce[0]);
        Debug.Log(Result_Announce[1]);
        Debug.Log(Result_Announce[2]);

        return Result_Announce;
    }



    private void now_Loading()
    {
        switch (counter%4)
        {
            case 0:
                text.text = "Now Loading";
                break;
            case 1:
                text.text = "Now Loading.";
                break;
            case 2:
                text.text = "Now Loading..";
                break;
            case 3:
                text.text = "Now Loading...";
                break;
            default:
                text.text = "Error";
                break;
        }

        counter++;
        if (counter > 4) counter = 0;
    }

    private void TriggerHand(Animator animation, string hand)
    {
        if (hand == "paper")
        {
            animation.SetTrigger("Paper");
        }
        else if (hand == "scissors")
        {
            animation.SetTrigger("Scissors");
        }
        else if (hand == "rock")
        {
            animation.SetTrigger("Rock");
        }
        else
        {
            text.text = "Error";
        }
    }

    // Start is called before the first frame update
    void Start()
    {
        time = 0;
        text = subtitle.GetComponent<Text>();
        myanimator = myHand.GetComponent<Animator>();
        enemyanimator = enemyHand.GetComponent<Animator>();

        audioSource = gameObject.GetComponent<AudioSource>();
        audioSource.Play();
        SceneManager.sceneLoaded += GameSceneLoaded;

        Debug.Log(player_status);
#if UNITY_EDITOR
        json_text = (Resources.Load("Hands") as TextAsset).text;
#endif
    }

    // Update is called once per frame
    void Update()
    {
        if (json_text == "")
        {
            now_Loading();
            return;
        }

        if (Responce_data == null)
        {
            audioSource.Stop();
            audioSource.clip = junksound;
            audioSource.Play();

            Responce_data = DeckInAnnounce_func(json_text);
        }


        if (time > 6.0f)
        {
            text.text = "";
            if (Responce_data[2] == "draw")
            {
                SceneManager.LoadScene("Hands_Animation_Scene");
            }
            else
            {
                SceneManager.LoadScene("Junk_Result");
            }
        }
        else if (time > 2.0f)
        {
            /*
             * TriggerHand(myanimator, hand);
             * TriggerHand(enemyanimator, hand);
             */
            TriggerHand(myanimator, Responce_data[0]);
            TriggerHand(enemyanimator, Responce_data[1]);
            text.text = "";
        }

        time += Time.deltaTime;
    }
}
